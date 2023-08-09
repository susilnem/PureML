import { fetchPublicProfile } from "../api/auth.server";

export async function load({ params }: any) {
  const username = params.username;
  return {
    publicProfile: await fetchPublicProfile(username),
  };
}
