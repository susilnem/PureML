<script>
  import { page } from "$app/stores";
  import error404 from "$lib/images/error/ErrorNotFound404.gif";
  import errorFunctional from "$lib/images/error/ErrorFunction.gif";
</script>

{#if $page.status === 404}
  <div class="bg-slate-50 text-slate-600 font-medium">
    <div
      class="flex px-12 pt-48 pb-12 md:justify-center items-center bg-slate-50"
    >
      <img src={error404} alt="Error" width="300" height="200" />
    </div>
  </div>
{:else}
  <div class="flex flex-col h-screen justify-center items-center bg-slate-50">
    <div class="text-3xl text-slate-600 font-medium">Oops!!</div>
    <div class="text-3xl text-slate-600 font-medium">
      Something went wrong :(
    </div>
    <img src={errorFunctional} alt="Error" width="500" />
  </div>
{/if}
