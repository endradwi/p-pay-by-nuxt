<template>
  <div>
    <div class="inline-block md:flex h-full md:h-screen w-full bg-blue-500">
      <div
        class="md:py-12 py-20 px-5 md:px-56 rounded-none md:rounded-r-[100px] w-full md:w-[1000px] bg-white flex flex-col gap-2"
      >
        <div class="items-center">
          <img src="~/assets/ppay_logo2.png" alt="" class="w-32 h-10" />
        </div>
        <div class="md:text-3xl text-lg font-medium">
          Start Accessing Banking Needs With All Devices and All Platforms With
          30.000+ Users
        </div>
        <div class="md:text-base text-xs">
          Transfering money is eassier than ever, you can access Zwallet
          wherever you are. Desktop, laptop, mobile phone? we cover all of that
          for you!
        </div>
        <div
          class=" flex flex-row md:flex-col justify-between md:justify-center w-full items-center gap-3"
        >
          <div
            class="btn border-slate-300 rounded-full  h-10 border py-3.5 px-20 md:px-44 flex items-center"
          >
            <FcGoogle />
            <span class="md:block hidden ">Sign In With Google</span>
          </div>
          <div
            class="btn border-slate-300 rounded-full  h-10 border py-3.5 px-20 md:px-[10.5rem] flex items-center"
          >
            <FaFacebook class="text-blue-700" />
            <span class="md:block hidden ">Sign In With Facebook</span>
          </div>
        </div>
        <div class="flex justify-between items-center gap-5 w-full">
          <div>
            <hr class="w-40 md:w-[250px] h-0.5 bg-info" />
          </div>
          <div class="text-info">or</div>
          <div>
            <hr class="w-40 md:w-[250px] h-0.5 bg-info" />
          </div>
        </div>
        <form @submit="" class="w-full flex flex-col gap-5">
          <label htmlFor="" class="flex flex-col gap-1">
            <div>
              <span class="text-base font-medium">Email</span>
            </div>
            <div
              class="input border-info w-full h-11 pl-5 flex items-center gap-3"
            >
              <MdOutlineMail class="text-info" />
              <input
                type="text"
                placeholder="Enter Your Email"
                class="w-full box-border"
              />
        </div>
          </label>
          <label htmlFor="" class="flex flex-col gap-1">
            <span class="text-base font-medium">Password</span>
            <div
              class="w-full h-11 border-black pl-5 flex items-center justify-between"
            >
              <div class="flex items-center gap-3">
                <MdOutlineVpnKey class="text-info" />
                <input
                  type="text"
                  placeholder="Enter Your Password"
                  class="w-full box-border"
                />
              </div>
              <!-- <div onClick={() => setShow(!isShow)}>
                    {isShow ? (
                      <IoEyeOutline class="text-info" />
                    ) : (
                      <IoEyeOffSharp class="text-info" />
                    )}
                  </div> -->
            </div>
          </label>
          <label htmlFor="" class="flex flex-col gap-1">
            <span class="text-base font-medium">Confirm Password</span>
            <div
              class="w-full h-11 input border-info pl-5 flex items-center justify-between"
            >
              <div class="flex items-center gap-3">
                <MdOutlineVpnKey class="text-info" />
                <input
                  type="text"
                  placeholder="Enter Your Password Again"
                  class="w-full box-border"
                />
              </div>
            </div>
          </label>
          <button class="bg-blue-500 py-4 rounded-xl text-white">Register</button>
        </form>
        <div class="w-full flex justify-center">
          Have An Account?
          <span class="text-blue-500">
            <NuxtLink to="/login">Login</NuxtLink>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>
