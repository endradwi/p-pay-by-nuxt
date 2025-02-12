<template>
     <div className="bg-blue-500 pt-16 pb-9 md:px-12 px-6 xl:px-32 flex gap-7 flex-col">
      <div className="border-b-2 border-b-slate-300 pb-24 gap-8 flex flex-col xl:flex-row justify-between">
        <div className="flex flex-col gap-6 max-w-64">
          <div className="w-56">
            <img src="~/assets/ppay.svg" alt="ppay" />
          </div>
          <div className="text-white">
            Clarity gives you the blocks and components you need to create a
            truly professional website.
          </div>
        </div>
        <div className="flex flex-col gap-6">
          <div className="text-lg font-semibold text-white">GET IN TOUCH</div>
          <div className="flex flex-col gap-3">
            <div className="text-white flex items-center gap-4">
              <!-- <BsTelephone /> -->
              <div>+62 5637 8882 9901</div>
            </div>
            <div className="text-white flex items-center gap-4">
              <!-- <FiMail /> -->
              <div>contact@zwallet.com</div>
            </div>
          </div>
        </div>
        <div className="flex flex-col gap-6">
          <div className="text-lg font-semibold text-white">SOCIAL MEDIA</div>
          <div className="flex gap-5">
            <div className="w-11 aspect-square bg-white rounded-full flex justify-center items-center text-xl text-blue-500 cursor-pointer">
              <!-- <FaTwitter /> -->
            </div>
            <div className="w-11 aspect-square bg-white rounded-full flex justify-center items-center text-xl text-blue-500 cursor-pointer">
              <!-- <FaFacebookF /> -->
            </div>
            <div className="w-11 aspect-square bg-white rounded-full flex justify-center items-center text-xl text-blue-500 cursor-pointer">
              <!-- <FaInstagram /> -->
            </div>
            <div className="w-11 aspect-square bg-white rounded-full flex justify-center items-center text-xl text-blue-500 cursor-pointer">
              <!-- <FaGithub /> -->
            </div>
          </div>
        </div>
        <form className="flex flex-col gap-6">
          <label
            htmlFor="subscriber"
            className="text-lg font-semibold text-white"
          >
            NEWSLETTER
          </label>
          <div className="flex flex-col gap-4">
            <div className="w-full lg:w-72 h-11 rounded-lg overflow-hidden relative">
              <input
                type="email"
                id="subscriber"
                name="subscriber"
                placeholder="Enter your email"
                className="outline-none w-full h-full box-border px-10"
              />
              <!-- <FiMail className="absolute top-[13px] text-lg left-3 text-gray-400" /> -->
            </div>
            <button className="btn btn-white lg:w-72 w-full h-9 font-light text-blue-500">
              Subscribe
            </button>
          </div>
        </form>
      </div>
      <div className="text-center text-white text-sm">
        © Copyright 2022, All Rights Reserved by ClarityUI
      </div>
    </div>
</template>