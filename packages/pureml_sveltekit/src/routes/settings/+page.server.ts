import { updateProfile } from "../api/auth.server.js";

export const actions = {
  default: async ({ cookies, request }: any) => {
    const form = await request.formData();
    const name = form.get("name");
    const bio = form.get("bio");
    const accesstoken = cookies.get("accessToken");
    const data = await updateProfile(name, bio, accesstoken);
    return data;
  },
};
