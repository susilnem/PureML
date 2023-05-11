import {
  fetchModelBranch,
  fetchModelVersions,
} from "../../../../../../api/models.server";

export async function load({ params, cookies }: any) {
  if (cookies.get("accessToken")) {
    const orgid = cookies.get("orgId");
    const modelid = params.modelId;
    const branchid = params.branchId;
    const accesstoken = cookies.get("accessToken");
    return {
      allBranches: await fetchModelBranch(orgid, modelid, accesstoken),
      versions: await fetchModelVersions(orgid, modelid, branchid, accesstoken),
    };
  }
}
