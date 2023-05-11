import { createAPIToken, deleteAPIToken } from "../../api/auth.server.js";

export const actions = {
  default: async ({ cookies, request }: any) => {
    const form = await request.formData();
    const createToken = form.get("createToken");
    const deleteToken = form.get("deleteToken");
    console.log("formA=", createToken, deleteToken);
    const accesstoken = cookies.get("accessToken");
    let newToken;
    let removeToken;
    if (createToken) {
      console.log("c");
      newToken = await createAPIToken(accesstoken);
      console.log("n=", newToken);
    } else newToken = null;
    if (deleteToken) {
      console.log("d");
      removeToken = await deleteAPIToken(deleteToken, accesstoken);
    } else removeToken = null;
    return { newToken, removeToken };
  },
};
