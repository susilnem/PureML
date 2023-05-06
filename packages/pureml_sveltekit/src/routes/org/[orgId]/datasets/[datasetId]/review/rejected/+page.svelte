<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import Tabbar from "$lib/components/Tabbar.svelte";
</script>

<Tabbar tabType="tertiary" tabContent="datasetReviewTab" tab="rejected" />
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
      <button type="submit" class="pb-6 w-full">
        <div
          class="hover:bg-slate-100 rounded-2xl flex justify-between items-center p-4"
        >
          <div class="flex items-center">
            <Avatar>
              <!-- {review.created_by.name
                                          .charAt(0)
                                          .toUpperCase()} -->
              K
            </Avatar>
            <div class="text-left text-sm text-slate-600 px-4">
              <a href="/" class="font-medium text-slate-800">
                <!-- href={`/${review.created_by.handle}`} -->
                "review.created_by.handle"
              </a>{" "}
              submitted{" "}
              "review.from_branch_version.version" of{" "}
              <a href="/" class="font-medium text-slate-800">
                <!-- href={`/org/${reviewData.params.orgId}/models/${review.model.name}`} -->
                "review.model.name"
              </a>{" "}
              from "review.from_branch.name" to{" "}
              "review.to_branch.name"
            </div>
          </div>
        </div>
      </button>
      <button type="submit" class="pb-6 w-full">
        <div
          class="hover:bg-slate-100 rounded-2xl flex justify-between items-center p-4"
        >
          <div class="flex items-center">
            <Avatar>
              <!-- {review.created_by.name
                                          .charAt(0)
                                          .toUpperCase()} -->
              K
            </Avatar>
            <div class="text-left text-sm text-slate-600 px-4">
              <a href="/" class="font-medium text-slate-800">
                <!-- href={`/${review.created_by.handle}`} -->
                "review.created_by.handle"
              </a>{" "}
              submitted{" "}
              "review.from_branch_version.version" of{" "}
              <a href="/" class="font-medium text-slate-800">
                <!-- href={`/org/${reviewData.params.orgId}/models/${review.model.name}`} -->
                "review.model.name"
              </a>{" "}
              from "review.from_branch.name" to{" "}
              "review.to_branch.name"
            </div>
          </div>
        </div>
      </button>
    </form>
  </div>
</div>
