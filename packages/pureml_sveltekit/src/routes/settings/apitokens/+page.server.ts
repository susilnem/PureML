import { createAPIToken, deleteAPIToken } from "../../api/auth.server.js";

export const actions = {
  default: async ({ cookies, request }: any) => {
    const form = await request.formData();
    const createToken = form.get("createToken");
    const deleteToken = form.get("deleteToken");
    const accesstoken = cookies.get("accessToken");
    let newToken;
    let removeToken;
    if (createToken) {
      newToken = await createAPIToken(accesstoken);
    } else newToken = null;
    if (deleteToken) {
      removeToken = await deleteAPIToken(deleteToken, accesstoken);
    } else removeToken = null;
    return { newToken, removeToken };
  },
};
