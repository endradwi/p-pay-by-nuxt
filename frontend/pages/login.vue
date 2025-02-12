<script lang="ts" setup>
useSeoMeta({
    title: 'Login Page'
})
</script>
<template>
  <div class="md:flex block h-screen bg-blue-500">
    <div
      class="md:py-28 md:px-36 px-5 py-40 rounded-none md:rounded-r-[100px] w-full h-full md:w-[1000px] bg-white flex flex-col gap-2"
    >
      <div class="items-center">
        <img src="~/assets/ppay_logo2.png" alt="ppay" class="w-32 h-10" />
      </div>
      <div class="text-3xl font-medium">Hello Welcome Back 👋</div>
      <div class="text-base">
        Fill out the form correctly or you can login with several option.
      </div>
      <div
        class="w-full flex flex-row md:flex-col justify-center items-center gap-3"
      >
        <div
          class="btn border-black rounded-full w-1/3 md:w-full h-10 border flex justify-between md:justify-center items-center"
        >
          <button
            class="w-full text-lg font-medium flex justify-center items-center gap-3 text-black"
          >
            <!-- <FcGoogle /> -->
            <span class="md:block hidden ">Sign In With Google</span>
          </button>
        </div>
        <div
          class="btn border-black rounded-full w-1/3 md:w-full h-10 border flex justify-between md:justify-center items-center"
        >
          <button
            class="w-full text-lg font-medium flex justify-center items-center gap-3 text-black"
          >
            <!-- <FaFacebook class="text-blue-700" /> -->
            <span class="md:block hidden ">Sign In With Facebook</span>
          </button>
        </div>
      </div>
      <div class="flex justify-between items-center gap-5">
        <div class="w-full">
          <hr class="md:w-full h-0.5 bg-black" />
        </div>
        <div class="text-black">or</div>
        <div class="w-full">
          <hr class="md:w-full h-0.5 bg-black" />
        </div>
      </div>
      <form
      class="w-full flex flex-col gap-2"
      >
      <label htmlFor="" class="flex flex-col gap-1">
        <div>
          <span class="text-base font-medium">Email</span>
        </div>
        <div
          class="input border w-full h-11 pl-5 flex items-center gap-3"
        >
          <MdOutlineMail class="text-black" />
          <input
            type="text"
            placeholder="Enter Your Email"
            class="w-full box-border"
          />
        </div>
      </label>
      <label  class="flex flex-col gap-1">
      <span class="text-base font-medium">Password</span>
      <div
        class="w-full h-11 input border-black pl-5 flex items-center justify-between"
      >
        <div class="flex items-center gap-3">
          <MdOutlineVpnKey class="text-black" />
          <input
            type="password"
            placeholder="Enter Your Password"
            class="w-full box-border"
          />
        </div>
        <!-- <div onClick={() => setShow(!isShow)}> -->
        <!-- {isShow ? ( -->
        <!-- <IoEyeOutline class="text-black" /> -->
        <!-- ) : ( -->
        <!-- <IoEyeOffSharp class="text-black" /> -->
        <!-- )} -->
      </div>
    </label>
      
    <NuxtLink to="/">
      <div class="w-full flex justify-end">
        <span class="link link-blue-500">Forgot Password?</span>
      </div>
    </NuxtLink>
    <button class="bg-blue-500 py-3 rounded-lg text-white">Login</button>
    <div class="w-full flex justify-center gap-2">
      Not Have An Account?
      <span class="link link-blue-500">
        <Link to="/register">Register</Link>
      </span>
    </div>
    </form>
    </div>
    
  </div>
  <!-- </div> -->
</template>
