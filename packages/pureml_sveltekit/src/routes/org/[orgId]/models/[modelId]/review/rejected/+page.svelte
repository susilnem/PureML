<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";

  export let data;
</script>

<Tabbar tabType="tertiary" tabContent="modelReviewTab" tab="rejected" />
<div class="px-12 pt-8 w-full h-[75%] overflow-auto">
  <div class="w-2/3">
    <form method="post">
      <input hidden readOnly name="reviewId" value="review.uuid" />
      <input
        hidden
        readOnly
        name="version"
        value="review.from_branch_version.version"
      />
      <input
        hidden
        readOnly
        name="fromBranch"
        value="review.from_branch.name"
      />
      <input hidden readOnly name="toBranch" value="review.to_branch.name" />
      <input hidden readOnly name="modelName" value="review.model.name" />
      {#if data.allCommits}
        {#each data.allCommits as reviewCommit}
          {#if reviewCommit.is_complete && !reviewCommit.is_accepted}
            <a
              href={`/org/${data.orgDetails[0].name}/models/${reviewCommit.model.name}/review/${reviewCommit.uuid}/rejected`}
              type="submit"
              class="pb-6 w-full"
            >
              <div
                class="hover:bg-slate-100 rounded-2xl flex justify-between items-center p-4"
              >
                <div class="flex items-center">
                  <Avatar>
                    {reviewCommit.created_by.name.charAt(0).toUpperCase()}
                  </Avatar>
                  <div class="text-left text-sm text-slate-600 px-4">
                    <a
                      href={`/${data.userDetails[0].handle}`}
                      class="font-medium text-slate-800"
                    >
                      <!-- href={`/${review.created_by.handle}`} -->
                      {reviewCommit.created_by.handle}
                    </a>{" "}
                    submitted{" "}
                    {reviewCommit.from_branch_version.version} of{" "}
                    <a
                      href={`/org/${data.orgDetails[0].name}/models/${reviewCommit.model.name}`}
                      class="font-medium text-slate-800"
                    >
                      {reviewCommit.model.name}
                    </a>{" "}
                    from {reviewCommit.from_branch.name} to{" "}
                    {reviewCommit.to_branch.name}
                  </div>
                </div>
              </div>
            </a>
          {/if}
        {/each}
      {/if}
    </form>
  </div>
</div>
