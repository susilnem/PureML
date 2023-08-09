import { updateOrg } from "../../api/org.server.js";

export const actions = {
  default: async ({ cookies, request }: any) => {
    const form = await request.formData();
    const orgname = form.get("orgname");
    const orgdesc = form.get("orgdesc");
    const orgid = form.get("orgid");
    const accesstoken = cookies.get("accessToken");
    const data = await updateOrg(orgdesc, orgname, orgid, accesstoken);
    return data;
  },
};
