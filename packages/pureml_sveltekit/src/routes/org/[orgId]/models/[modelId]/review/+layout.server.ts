import { fetchModelReview } from "../../../../../api/models.server";

export async function load({ params, cookies }: any) {
  if (cookies.get("accessToken")) {
    const orgid = cookies.get("orgId");
    const modelName = params.modelId;
    const accesstoken = cookies.get("accessToken");
    return {
      allCommits: await fetchModelReview(orgid, modelName, accesstoken),
    };
  }
}
