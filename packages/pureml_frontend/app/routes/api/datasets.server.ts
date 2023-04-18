const backendUrl = process.env.BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;

// ###########################################################################

// ########################### dataset details api ###########################

export async function fetchDatasets(orgId: string, accessToken: string) {
  const url = makeUrl(`org/${orgId}/dataset/all`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}

export async function fetchDatasetByName(
  orgId: string,
  datasetName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}

// ############################# dataset readme #############################

export async function fetchDatasetReadme(
  orgId: string,
  datasetName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/readme/version`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}

export async function writeDatasetReadme(
  orgId: string,
  datasetName: string,
  content: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/readme`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      content: content,
      file_type: "html",
    }),
  }).then((res) => res.json());
  return res;
}

// ######################### dataset branch details ###########################

export async function fetchDatasetBranch(
  orgId: string,
  datasetName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/branch`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  if (res.data !== null) {
    let b: Object[] = [];
    res.data.forEach((branch: any) => {
      b.push({ value: branch.name, label: branch.name });
    });
    return b;
  }
  return res;
}

// ########################## dataset version details ###########################

export async function fetchOneDatasetVersion(
  orgId: string,
  datasetId: string,
  branchName: string,
  version: string,
  accessToken: string
) {
  const url = makeUrl(
    `org/${orgId}/dataset/${datasetId}/branch/${branchName}/version/${version}`
  );

  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}

export async function fetchDatasetVersions(
  orgId: string,
  datasetName: string,
  branchName: string,
  accessToken: string
) {
  const url = makeUrl(
    `org/${orgId}/dataset/${datasetName}/branch/${branchName}/version`
  );
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}

// ######################### dataset review ###########################

export async function fetchDatasetReview(
  orgId: string,
  datasetName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/review`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}

export async function submitDatasetReview(
  orgId: string,
  datasetName: string,
  branch: string,
  version: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/review/create`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      description: "desc",
      from_branch: branch,
      from_branch_version: version,
      is_accepted: false,
      is_complete: false,
      title: "title",
      to_branch: "main",
    }),
  }).then((res) => res.json());
  return res;
}

export async function updateDatasetReview(
  orgId: string,
  datasetName: string,
  reviewId: string,
  accepted: string,
  accessToken: string
) {
  const url = makeUrl(
    `org/${orgId}/dataset/${datasetName}/review/${reviewId}/update`
  );
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      description: "-",
      is_accepted: accepted === "true",
      is_complete: true,
      title: "-",
    }),
  }).then((res) => res.json());
  return res;
}
