/** @type {import('./$types').Actions}  */

import type { Actions } from "@sveltejs/kit";
import { fetchSignIn } from "../../api/auth.server";
import { fetchAllOrgs } from "../../api/org.server";

export const actions: Actions = {
  default: async ({ cookies, request }: any) => {
    const form = await request.formData();
    const email = form.get("email");
    const password = form.get("password");
    const data = await fetchSignIn(email, password);

    if (data.status === 202 || data.status === 200) {
      const accesstoken = data.data[0].accessToken;
      cookies.set("accessToken", accesstoken, {
        path: "/",
        sameSite: "strict",
      });
      const org = await fetchAllOrgs(accesstoken);
      cookies.set("orgId", org[0].org.uuid, {
        path: "/",
        sameSite: "strict",
      });
      cookies.set("orgName", org[0].org.name, {
        path: "/",
        sameSite: "strict",
      });
    }
    return data;

    //   const session = await getSession(request.headers.get("Cookie"));
    //   const sessionId = new URL(request.url);
    //   const verifySessionId = sessionId.searchParams.get("sessionid");
    //   const accessToken = data.data[0].accessToken;
    //   session.set("accessToken", accessToken);
    //   if (!verifySessionId && data.message === "User logged in") {
    //     return json(data, {
    //       headers: {
    //         "Set-Cookie": await commitSession(session),
    //       },
    //     });
    //   } else if (verifySessionId && data.message === "User logged in") {
    //     if (verifySessionId && accessToken) {
    //       const orgId = session.get("orgId");
    //       if (!orgId) return null;
    //       const profile = await fetchUserSettings(accessToken);
    //       const orgDetails = await fetchOrgDetails(
    //         orgId,
    //         session.get("accessToken")
    //       );
    //       const verifySession = await fetchVerifySession(
    //         accessToken,
    //         verifySessionId
    //       );
    //       return json(
    //         { profile, orgDetails, verifySession, verifySessionId },
    //         {
    //           headers: {
    //             "Set-Cookie": await commitSession(session),
    //           },
    //         }
    //       );
    //     }
    //   }
    //   return { data, verifySessionId, ok: true };
  },
};
