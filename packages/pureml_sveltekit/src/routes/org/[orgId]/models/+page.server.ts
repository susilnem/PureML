import { redirect } from "@sveltejs/kit";

export function load({ params }) {
  const orgId = params.orgId;
  throw redirect(307, `/org/${orgId}`);
}
