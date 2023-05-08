import { redirect } from "@sveltejs/kit";
import type { Actions } from "./$types";

export const actions: Actions = {
  default({ cookies }) {
    cookies.delete("accessToken", { path: "/" });
    throw redirect(302, "/auth/signin");
  },
};
