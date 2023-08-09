import { Suspense, useEffect, useState } from "react"
import Tabbar from "~/components/Tabbar"
import Button from "~/components/ui/Button"
import Input from "~/components/ui/Input"
import Loader from "~/components/ui/Loading"
import { getSession } from "~/session"
import { Form, useActionData, useLoaderData } from "@remix-run/react"
import { toast } from "react-toastify"
import Modal from "~/components/ui/Modal"
import { Copy } from "lucide-react"
import {
	createAPIToken,
	deleteAPIToken,
	fetchAPITokens,
} from "~/routes/api/auth.server"
import { ScrollUpButton } from "@radix-ui/react-select"

export async function loader({ request }: any) {
	const session = await getSession(request.headers.get("Cookie"))
	const accesstoken = session.get("accessToken")
	const listTokens = await fetchAPITokens(accesstoken)
	return listTokens
}

export const action = async ({ request }: any) => {
	const form = await request.formData()
	const createToken = form.get("createToken")
	const tokenName = form.get("name")
	const deleteToken = form.get("deleteToken")
	console.log("form=", createToken, tokenName, deleteToken)
	const session = await getSession(request.headers.get("Cookie"))
	const accessToken = session.get("accessToken")
	let newToken
	let removeToken
	if (createToken) {
		console.log("c")
		newToken = await createAPIToken(accessToken, tokenName)
		console.log("n=", newToken)
	} else newToken = null
	if (deleteToken) {
		console.log("d")
		removeToken = await deleteAPIToken(deleteToken, accessToken)
	} else removeToken = null
	return { newToken, removeToken }
}

export default function APITokensSetting() {
	const data = useLoaderData()
	// console.log("data=", data)
	const actionData = useActionData()
	// console.log("actD=", actionData)
	useEffect(() => {
		if (!actionData) return
		if (actionData.removeToken) {
			if (actionData.removeToken.message === "Token deleted")
				toast.success("API Key ID deleted successfully!")
			else toast.error("Something went wrong!")
		}
	}, [actionData])

	const [validName, setValidName] = useState(false)

	function copyOrgId(text: string) {
		navigator.clipboard.writeText(text)
		toast.success("Copied to clipboard!")
	}

	return (
		<Suspense fallback={<Loader />}>
			<div className='flex justify-center w-full border-b-2 border-slate-100'>
				<div className='w-full 2xl:max-w-screen-2xl'>
					<Tabbar intent='primarySettingTab' tab='apiToken' />
				</div>
			</div>
			<div className='flex justify-center w-full'>
				<div className='bg-slate-50 h-screen overflow-hidden w-full 2xl:max-w-screen-2xl'>
					<div className='flex flex-col gap-y-8 pt-8 pb-12 px-12 w-full h-4/5 overflow-auto rounded-tl-lg rounded-tr-lg'>
						<div className='text-slate-800'>API Tokens</div>
						<table className='table w-2/5'>
							<thead>
								<tr className='bg-slate-100 border border-slate-200 rounded-tl-lg rounded-tr-lg'>
									<th className='text-slate-950 font-medium'>API Token Name</th>
									<th className='text-slate-950 font-medium'>Created</th>
									<th className='text-slate-950 font-medium'>Actions</th>
								</tr>
							</thead>
							<tbody>
								{data ? (
									<>
										{data.map((token: any, index: number) => (
											<tr
												className='border border-slate-200 hover:bg-slate-50'
												key={index}>
												<td className='text-lg'>{token.name}</td>
												<td className='text-lg'>
													{token.created_at
														.replace("T", " ")
														.replace("Z", "")
														.slice(0, 19)}
												</td>
												<td className='text-lg'>
													<Form method='post' reloadDocument className='w-full'>
														<input
															className='hidden'
															type='text'
															name='deleteToken'
															defaultValue={token.uuid}
														/>
														<button className='text-lg p-2 hover:bg-slate-100 hover:rounded-lg'>
															Delete
														</button>
													</Form>
												</td>
											</tr>
										))}
									</>
								) : (
									<div>No API keys found</div>
								)}
							</tbody>
						</table>
						<Form method='post'>
							<input
								className='hidden'
								type='text'
								name='createToken'
								defaultValue='createToken'
							/>
							<label htmlFor='name' className='font-medium'>
								Name
								<Input
									intent='primary'
									type='text'
									name='name'
									fullWidth={false}
									aria-label='name'
									data-testid='name'
									required
									onChange={(e) => {
										if (e.target.value == "") {
											setValidName(false)
										} else {
                      setValidName(true)
                    }
									}}
								/>
							</label>
							<div className='mt-5'>
								<Modal
									btnName={
										<Button disabled={!validName} intent='primary' fullWidth={false}>
											Create API Key
										</Button>
									}
									title="Token created">
									<div className='text-slate-600'>
										This is the last time you can see the API Key Secret! Sorry, security...
									</div>
									{actionData && actionData.newToken && (
										<div className='flex flex-col gap-y-6 pt-8'>
											<label htmlFor='keysecret' className='pb-1'>
												API Key Secret
												<div className='input-icons'>
													<input
														id='orgid'
														type='text'
														name='orgid'
														defaultValue={
															actionData.newToken.data[0].api_token_secret ||
															"API Key Secret"
														}
														value={actionData.newToken.data[0].api_token_secret}
														aria-label='orgid'
														data-testid='orgid'
														required
														disabled
														className='input-field rounded !w-full !pr-12 truncate'
													/>
													<i>
														<Copy
															className='text-slate-400 hover:text-slate-600 w-4 cursor-pointer'
															onClick={() =>
																copyOrgId(
																	actionData.newToken.data[0].api_token_secret
																)
															}
														/>
													</i>
												</div>
											</label>
										</div>
									)}
								</Modal>
							</div>
						</Form>
					</div>
				</div>
			</div>
		</Suspense>
	)
}

// ############################ error boundary ###########################

export function ErrorBoundary() {
	return (
		<div className='flex flex-col h-screen justify-center items-center bg-slate-50'>
			<div className='text-3xl text-slate-600 font-medium'>Oops!!</div>
			<div className='text-3xl text-slate-600 font-medium'>
				Something went wrong :(
			</div>
			<img src='/error/ErrorFunction.gif' alt='Error' width='500' />
		</div>
	)
}

export function CatchBoundary() {
	return (
		<div className='flex flex-col h-screen justify-center items-center bg-slate-50'>
			<div className='text-3xl text-slate-600 font-medium'>Oops!!</div>
			<div className='text-3xl text-slate-600 font-medium'>
				Something went wrong :(
			</div>
			<img src='/error/ErrorFunction.gif' alt='Error' width='500' />
		</div>
	)
}
