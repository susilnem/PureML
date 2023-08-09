import { fetchModelVersions } from "../../../../../../api/models.server";
import type { Actions } from "./$types";

export const actions = {
  changeBranch: async ({ cookies, params, request }: any) => {
    const form = await request.formData();
    const branch = form.get("branchSelected");
    const orgid = cookies.get("orgId");
    const modelid = params.modelId;
    const accesstoken = cookies.get("accessToken");
    const branchVersions = await fetchModelVersions(
      orgid,
      modelid,
      branch,
      accesstoken
    );
    if (branchVersions != null) {
      return { branchVersions };
    } else return { branchVersions: [{ version: { version: "-" } }] };
  },
} satisfies Actions;
