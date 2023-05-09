<script>
  import Avatar from "$lib/components/Avatar.svelte";
  import Card from "$lib/components/Card.svelte";
  import Error from "../+error.svelte";
  import { page } from "$app/stores";

  export let data;
</script>

<svelte:head>
  <title>{$page.params.username || "Profile"} | PureML</title>
  <meta name="description" content="Developer platform for Production ML." />
</svelte:head>

<!-- ##### public profile for logged in users #### -->
<div>
  {#if data.userDetails && data.publicProfile}
    {#if data.publicProfile[0]}
      <div class="flex flex-col w-screen h-screen overflow-hidden">
        <div class="flex justify-center w-full">
          <div
            class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
          >
            <div
              class="flex flex-col gap-y-8 md:flex md:flex-row px-12 pt-8 pb-12 text-slate-600 font-medium overflow-auto"
            >
              <div class="h-full md:w-28 md:w-44 lg:w-56 2xl:w-96">
                <div
                  class="h-28 w-28 md:h-36 md:w-36 lg:w-56 lg:h-56 2xl:h-96 2xl:w-96 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-lg"
                >
                  <Avatar intent="unfilled">
                    {data.publicProfile[0]?.name.charAt(0).toUpperCase() ||
                      "User"}
                  </Avatar>
                </div>
                <div class="pt-6 font-medium text-base text-slate-600">
                  {data.publicProfile[0]?.name || "Name"}
                </div>
                <div class="pb-6 text-base font-normal">
                  {data.publicProfile[0]?.email || "Email"}
                </div>
                <!-- {/* <Button aria-label="follow" intent="primary" >
                          Follow
                        </Button> */} -->
                <div class="text-base">
                  <span>Bio</span>
                </div>
                <div class="font-normal text-base text-slate-600">
                  {data.publicProfile[0]?.bio || "Add your bio"}
                </div>
                <!-- {/* <div class="text-base pt-8">Organizations</div>
                        {data.publicProfile[0].orgs.length !== 0 ? (
                          <div>
                            {data.publicProfile[0].orgs.map((org: any) => (
                              <Button intent="org"  key={org.name}>
                                {org.name}
                              </Button>
                            ))}
                          </div>
                        ) : (
                          "-"
                        )} */} -->
              </div>
              <div class="md:pl-12 w-full md:w-3/4 2xl:w-1/2">
                <!-- {/* <div>
                          <div class="flex justify-between">
                            <div class="text-slate-800 font-medium font-base">
                              Total contributions this year
                            </div>
                          </div>
                          <div class="pt-6">
                            <CalendarHeatmap
                              startDate={shiftDate(today, -270)}
                              endDate={today}
                              values={randomValues}
                              classForValue={(value: any) => {
                                if (!value) {
                                  return "color-empty";
                                }
                                return `color-github-${value.count}`;
                              }}
                              showWeekdayLabels={true}
                            />
                          </div>
                          <div class="flex justify-end"><img src="/imgs/ContributionScale.png" /></div>
                        </div> */} -->
                <div class="font-medium">Overview</div>
                <div
                  class="pt-8 flex flex-col gap-y-4 md:flex md:flex-row w-full md:gap-x-4"
                >
                  <div class="w-full">
                    <Card
                      title="Models"
                      description=""
                      count={data.publicProfile[0]?.number_of_models || "0"}
                    />
                  </div>
                  <div class="w-full">
                    <Card
                      title="Datasets"
                      description=""
                      count={data.publicProfile[0]?.number_of_datasets || "0"}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    {/if}
    <!-- ##### anonymous view of public profile #### -->
  {:else if data.publicProfile}
    {#if data.publicProfile[0]}
      <div class="flex flex-col w-screen h-screen overflow-hidden">
        <div class="flex justify-center w-full">
          <div
            class="bg-slate-50 flex flex-col h-screen overflow-hidden w-full 2xl:max-w-screen-2xl"
          >
            <div
              class="flex flex-col gap-y-8 md:flex md:flex-row px-12 pt-8 pb-12 text-slate-600 font-medium overflow-auto"
            >
              <div class="h-full md:w-28 md:w-44 lg:w-56 2xl:w-96">
                <div
                  class="h-28 w-28 md:h-36 md:w-36 lg:w-56 lg:h-56 2xl:h-96 2xl:w-96 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-lg"
                >
                  <Avatar intent="unfilled">
                    {data.publicProfile[0]?.name.charAt(0).toUpperCase() ||
                      "User"}
                  </Avatar>
                </div>
                <div class="pt-6 font-medium text-base text-slate-600">
                  {data.publicProfile[0]?.name || "Name"}
                </div>
                <div class="pb-6 text-base font-normal">
                  {data.publicProfile[0]?.email || "Email"}
                </div>
                <!-- <Button aria-label="follow" intent="primary" >
              Follow
            </Button> -->
                <div class="text-base">
                  <span>Bio</span>
                </div>
                <div class="font-medium text-base text-slate-600">
                  {data.publicProfile[0]?.bio || "Add your bio"}
                </div>
                <!-- <div class="text-base pt-8">Organizations</div>
            {data.publicProfile[0].orgs.length !== 0 ? (
              <div>
                {data.publicProfile[0].orgs.map((org: any) => (
                  <Button intent="org"  key={org.name}>
                    {org.name}
                  </Button>
                ))}
              </div>
            ) : (
              "-"
            )} -->
              </div>
              <div class="md:pl-12 w-full md:w-3/4 2xl:w-1/2">
                <!-- <div>
                <div class="flex justify-between">
                  <div class="text-slate-800 font-medium font-base">
                    Total contributions this year
                  </div>
                </div>
                <div class="pt-6">
                  <CalendarHeatmap
                    startDate={shiftDate(today, -270)}
                    endDate={today}
                    values={randomValues}
                    classForValue={(value: any) => {
                      if (!value) {
                        return "color-empty";
                      }
                      return `color-github-${value.count}`;
                    }}
                    showWeekdayLabels={true}
                  />
                </div>
                <div class="flex justify-end"><img src="/imgs/ContributionScale.png" /></div>
              </div> -->
                <div class="font-medium">Overview</div>
                <div
                  class="pt-8 flex flex-col gap-y-4 md:flex md:flex-row w-full md:gap-x-4"
                >
                  <div class="w-full">
                    <Card
                      title="Models"
                      description=""
                      count={data.publicProfile[0]?.number_of_models || "0"}
                    />
                  </div>
                  <div class="w-full">
                    <Card
                      title="Datasets"
                      description=""
                      count={data.publicProfile[0]?.number_of_datasets || "0"}
                    />
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    {/if}

    <!-- ##### when profile does not exist #### -->
  {:else}
    <Error />
  {/if}
</div>
