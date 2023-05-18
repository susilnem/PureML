import { fetchModelVersions } from "../../../../../../api/models.server";
import type { Actions } from "./$types";

export const actions = {
  changeBranch: async ({ cookies, params, request }: any) => {
    console.log("in");
    const form = await request.formData();
    const branch = form.get("branchSelected");
    console.log(branch);
    const orgid = cookies.get("orgId");
    const modelid = params.modelId;
    const accesstoken = cookies.get("accessToken");
    const branchVersions = await fetchModelVersions(
      orgid,
      modelid,
      branch,
      accesstoken
    );
    if (branchVersions.length > 0) {
      return { success: true, branchVersions };
    } else return { success: false };
  },
  changeVersion: async ({ request }: any) => {
    const form = await request.formData();
    const branch = form.get("branchSelected");
    const version = form.get("versionSelected");
    console.log(branch, version);
    return;
  },
} satisfies Actions;
