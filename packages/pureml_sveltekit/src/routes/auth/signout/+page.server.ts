import { redirect } from "@sveltejs/kit"
import type { Actions } from "./$types"
import { base } from "$app/paths"

export const actions: Actions = {
	default({ cookies }) {
		cookies.delete("accessToken", { path: "/" })
		cookies.delete("orgName", { path: "/" })
		cookies.delete("orgId", { path: "/" })
		throw redirect(302, base + "/auth/signin")
	},
}
