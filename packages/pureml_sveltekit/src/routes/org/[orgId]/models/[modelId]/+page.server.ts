import { fetchModelByName, fetchModels } from "../../../../api/models.server";

export async function load({ params, cookies }: any) {
  if (cookies.get("accessToken")) {
    const accesstoken = cookies.get("accessToken");
    const orgid = cookies.get("orgId");
    const modelid = params.modelId;
    return {
      models: await fetchModels(orgid, accesstoken),
      modelDetails: await fetchModelByName(orgid, modelid, accesstoken),
    };
  }
}
