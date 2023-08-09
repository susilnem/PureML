import { fetchAPITokens, updateProfile } from "../api/auth.server";

export async function load({ cookies }: any) {
  if (cookies.get("accessToken")) {
    const accesstoken = cookies.get("accessToken");
    return {
      apiTokens: await fetchAPITokens(accesstoken),
    };
  }
}
