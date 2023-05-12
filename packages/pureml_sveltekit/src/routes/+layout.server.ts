import { fetchUserSettings } from "./api/auth.server";
import { fetchModelBranch, fetchModelVersions } from "./api/models.server";
import { fetchOrgDetails } from "./api/org.server";

export async function load({ cookies }: any) {
  if (cookies.get("accessToken")) {
    const accesstoken = cookies.get("accessToken");
    const orgid = cookies.get("orgId");
    return {
      orgDetails: await fetchOrgDetails(orgid, accesstoken),
      userDetails: await fetchUserSettings(accesstoken),
    };
  }
}
