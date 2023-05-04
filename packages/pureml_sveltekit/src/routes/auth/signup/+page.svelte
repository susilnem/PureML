<script>
  import Button from "$lib/components/Button.svelte";
  import Input from "$lib/components/Input.svelte";
  import Link from "$lib/components/Link.svelte";
  import logo from "$lib/logos/PureMLLogoText.svg";
</script>

<div class="flex justify-center">
  <div class="w-fit text-center">
    <div class="flex justify-center items-center pb-16">
      <img src={logo} alt="Logo" class="w-36" />
    </div>
    <form method="post" class="text-slate-600 flex flex-col text-left">
      <div class="flex flex-col gap-y-12">
        <div class="flex flex-col gap-y-6">
          <label for="name" class="font-medium">
            <div class="text-slate-600 text-left">Name</div>
            <Input type="text" required intent="primary" fullWidth={false} />
          </label>
          <label for="username" class="font-medium">
            <div class="text-slate-600 text-left">Username</div>
            <Input type="text" required intent="primary" fullWidth={false} />
          </label>
          <label for="email" class="font-medium">
            <div class="text-slate-600 text-left">Email</div>
            <Input type="email" required intent="primary" fullWidth={false} />
          </label>
          <label for="password" class="font-medium">
            <div class="text-slate-600 text-left">Password</div>
            <Input
              type="password"
              required
              intent="primary"
              fullWidth={false}
            />
          </label>
        </div>
        <div class="w-[24rem] flex flex-col gap-y-6">
          <span class="text-slate-600">
            By clicking the Sign in button, you agree to PureML’s{" "}
            <a
              href="/TermsAndCondition.pdf"
              target="_blank"
              class="text-blue-250"
            >
              Terms of Service
            </a>{" "}
            and{" "}
            <a href="/PrivacyPolicy.pdf" target="_blank" class="text-blue-250">
              Privacy Policy
            </a>
            .
          </span>
          <Button intent="primary">Sign Up</Button>
        </div>
      </div>
    </form>
    <div class="flex items-center text-slate-600 justify-center mt-6">
      <Link intent="secondary" hyperlink="/auth/forgot_password">
        Forgot Password?
      </Link>
      <p class="px-2 text-slate-400">|</p>
      <div class="flex items-center space-x-1 font-medium">
        <!-- <span>Already a user?</span> -->
        <Link intent="secondary" hyperlink="/auth/signin">Sign In</Link>
      </div>
    </div>
  </div>
</div>
