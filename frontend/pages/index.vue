<script lang="ts" setup>
useSeoMeta({
    title: 'Home Page'
})
</script>
<template>
    <div>
        <Navbar />
      <main>
        <div class="py-36 flex px-6 gap-6 lg:gap-0 md:px-12 items-center lg:px-32 flex-col-reverse lg:flex-row">
          <div class="flex flex-col gap-6 lg:w-1/2">
            <div class=" text-4xl lg:text-5xl text-[#0B132A] font-medium lg:max-w-3xl">
              Smart Way to Your Financial Business
            </div>
            <div class="max-w-md">
              <span class="text-base text-info font-light">
                We bring you a mobile app for banking problems that oftenly
                wasting much of your times.
              </span>
            </div>
            <NuxtLink to="/"><div>
              <button class="bg-blue-500 py-3 px-5 rounded-lg font-light px-8 text-white">
                Get Started
              </button>
            </div></NuxtLink>
            <div>
              <span>Available on</span>
            </div>
            <div class="flex gap-6">
              <img src="~/assets/gplay.svg" alt="gplay" />
              <img src="~/assets/appstore.svg" alt="appstore" />
            </div>
          </div>
          <div class="flex shrink-0 w-1/2 justify-center">
            <img src="~/assets/home-page.png" alt="cellphone" />
          </div>
        </div>
        <div class="bg-abuMuda gap-20 flex-wrap flex justify-center py-9 px-32">
          <img src="~/assets/microsoft.svg" alt="microsoft" />
          <img src="~/assets/dropbox.svg" alt="dropbox" />
          <img src="~/assets/hnm.svg" alt="hnm" />
          <img src="~/assets/airbnb.svg" alt="airbnb" />
          <img src="~/assets/canon.svg" alt="canon" />
          <img src="~/assets/dell.svg" alt="dell" />
        </div>
        <div class="py-36 lg:px-12 px-6 md:px-32 gap-10 xl:flex-row flex-col text-center lg:text-start flex justify-around items-center">
          <div class="flex flex-col gap-6">
            <div>
              <span class="text-4xl font-medium">About The Aplication</span>
            </div>
            <div class="max-w-lg">
              <span class="text-base text-info">
                We have some great features from the application and it’s
                totally free to use by all users around the world.
              </span>
            </div>
          </div>
          <div class="flex flex-col lg:flex-row gap-6">
            <div class="rounded-lg bg-blue-500 w-full lg:w-52 flex flex-col gap-5 py-9 text-center px-3 items-center">
              <div class="w-14 aspect-square bg-white rounded-full flex justify-center items-center text-3xl text-blue-500">
                <!-- <CiHeadphones /> -->
              </div>
              <div class="text-lg font-bold text-white">24/7 Support</div>
              <div class="text-base text-white">
                We have 24/7 contact support so you can contact us whenever you
                want and we will respond it.
              </div>
            </div>
            <div class="rounded-lg bg-blue-500 w-full lg:w-52 flex flex-col gap-5 py-9 text-center px-3 items-center">
              <div class="w-14 aspect-square bg-white rounded-full flex justify-center items-center text-3xl text-blue-500">
                <!-- <IoShieldCheckmarkOutline /> -->
              </div>
              <div class="text-lg font-bold text-white">Data Privacy</div>
              <div class="text-base text-white">
                We make sure your data is safe in our database and we will
                encrypt any data you submitted to us.
              </div>
            </div>
            <div class="rounded-lg bg-blue-500 w-full lg:w-52 flex flex-col gap-5 py-9 text-center px-3 items-center">
              <div class="w-14 aspect-square bg-white rounded-full flex justify-center items-center text-3xl text-blue-500">
                <!-- <IoDownloadOutline /> -->
              </div>
              <div class="text-lg font-bold text-white">
                Easy Download
              </div>
              <div class="text-base text-white">
                Zwallet is 100% totally free to use it’s now available on Google
                Play Store and App Store.
              </div>
            </div>
          </div>
        </div>
        <div class="bg-blue-500 flex flex-col xl:flex-row items-center justify-center xl:justify-center xl:items-end py-24 md:px-12 px-6 lg:px-32 gap-56">
          <div class="flex justify-center items-center">
            <img src="~/assets/Mobile-dashboard.png" alt="" />
          </div>
          <div class="flex flex-col gap-6">
            <div class="text-white font-medium text-4xl">
              All The Great Zwallet Features.
            </div>
            <div class="text-base text-white font-light">
              We have some great features from the application and it’s totally
              free to use by all users around the world.
            </div>
            <div class="flex gap-3 items-center lg:items-start">
              <div class="w-14 flex-shrink-0 h-14 bg-white rounded-full flex justify-center items-center text-3xl text-blue-500">
                <!-- <PiMoneyWavy /> -->
              </div>
              <div class="flex text-white flex-col gap-2 justify-center">
                <div class="font-semibold text-lg">Small Fee</div>
                <div class="text-base font-light">
                  We only charge 5% of every success transaction done in Zwallet
                  app.
                </div>
              </div>
            </div>
            <div class="flex gap-3 items-center lg:items-start">
              <div class="w-14 flex-shrink-0 h-14 bg-white rounded-full flex justify-center items-center text-3xl text-blue-500">
                <!-- <IoShieldCheckmarkOutline /> -->
              </div>
              <div class="flex text-white flex-col gap-2 justify-center">
                <div class="font-semibold text-lg">Data Secured</div>
                <div class="text-base font-light">
                  All your data is secured properly in our system and it’s
                  encrypted.
                </div>
              </div>
            </div>
            <div class="flex gap-3 items-center lg:items-start">
              <div class="w-14 flex-shrink-0 h-14 bg-white rounded-full flex justify-center items-center text-3xl text-blue-500">
                <LuUserCheck />
              </div>
              <div class="flex text-white flex-col gap-2 justify-center">
                <div class="font-semibold text-lg">User Friendly</div>
                <div class="text-base font-light">
                  Zwallet come up with modern and sleek design and not
                  complicated.
                </div>
              </div>
            </div>
            <div>
             <NavLink to="/dashboard" class="btn w-full md:w-44 btn-white font-light text-blue-500 px-11"> <button >
                Get Started
              </button> </NavLink>
            </div>
          </div>
        </div>
        <div class="md:px-12 px-6 lg:px-32 flex flex-col items-center gap-8 lg:gap-20 text-center py-20">
          <div class="text-3xl  max-w-96 lg:max-w-full  lg:text-5xl font-medium">
            Here From Our Customer
          </div>
          <div class="text-info">
            We always do our best for our customers to stay comfortable using
            the applications we provide
          </div>
          <div class="flex gap-11 items-center justify-center">
            <div class="btn hidden xl:flex hover:bg-blue-500 hover:text-white rounded-full w-12 aspect-square border-none bg-abuMuda">
              <FaArrowLeftLong />
            </div>
            <div class="rounded-lg bg-abuMuda w-72 hidden xl:flex flex-col gap-5 py-14 text-[#0B132A] text-center px-3 items-center">
              <div class="avatar placeholder">
                <div class="w-12 h-12 rounded-full">
                  <img src="~/assets/avatar.png" alt="avatar" />
                </div>
              </div>
              <div class="text-lg font-bold">Sherina Claw</div>
              <div class="flex gap-4 items-center">
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <div class="text-info">5.0</div>
              </div>
              <div class="text-5xl font-bold">“</div>
              <div class="text-base text-info">
                “I use Zwallet to manage all financial needs. It’s super easy to
                use and it’s 100% free app”
              </div>
            </div>
            <div class="rounded-lg bg-abuMuda w-72 flex flex-col gap-5 py-14 text-[#0B132A] text-center px-3 items-center">
              <div class="avatar placeholder">
                <div class="w-12 h-12 rounded-full">
                    <img src="~/assets/avatar.png" alt="avatar"  />
                </div>
              </div>
              <div class="text-lg font-bold">James Bond</div>
              <div class="flex gap-4 items-center">
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <div class="text-info">5.0</div>
              </div>
              <div class="text-5xl font-bold">“</div>
              <div class="text-base text-info">
                “Since I’m using this app, I’m not going to move to another
                similar app. Thank you Zwallet!”
              </div>
            </div>
            <div class="rounded-lg bg-abuMuda w-72 hidden xl:flex flex-col gap-5 py-14 text-[#0B132A] text-center px-3 items-center">
              <div class="avatar placeholder">
                <div class="w-12 h-12 rounded-full">
                    <img src="~/assets/avatar.png" alt="avatar" />
                </div>
              </div>
              <div class="text-lg font-bold">Ujang Kayu</div>
              <div class="flex gap-4 items-center">
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <img src="~/assets/star.svg" alt="star" />
                <div class="text-info">5.0</div>
              </div>
              <div class="text-5xl font-bold">“</div>
              <div class="text-base text-info">
                “I use this app since 2 years ago and this is the best app that
                I’ve ever use in my entire life”
              </div>
            </div>
            <div class="btn hidden xl:flex hover:bg-blue-500 hover:text-white rounded-full w-12 aspect-square border-none bg-abuMuda">
              <!-- <FaArrowRightLong /> -->
            </div>
          </div>
          <div class="flex gap-3 justify-center">
            <div class="bg-blue-500 w-6 rounded-full"></div>
            <div class="bg-slate-200 w-2 aspect-square rounded-full"></div>
            <div class="bg-slate-200 w-2 aspect-square rounded-full"></div>
            <div class="bg-slate-200 w-2 aspect-square rounded-full"></div>
          </div>
        </div>
      </main>
        <Footer></Footer>
    </div>
</template>