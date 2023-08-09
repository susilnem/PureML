import { redirect } from "@sveltejs/kit";

export function load({ params }: any) {
  const orgId = params.orgId;
  throw redirect(307, `/org/${orgId}`);
}
