<template>
  <div>
    <div
      className="flex navbar justify-between bg-blue-500 px-5 md:px-12 lg:px-32 py-4 shadow-md"
    >
      <img className="h-8" src="~/assets/ppay.svg" alt="ppay" />
      <div className="md:flex gap-6 hidden">
        <NuxtLink to="/login">
          <button
            className="py-2 px-5 rounded-lg text-white bg-blue-500 border border-white text-sm font-normal"
          >
            Sign In
          </button>
        </NuxtLink>
        <NuxtLink to="/register">
          <button
            className="py-2 px-5 rounded-lg bg-white border-none text-blue-500 text-sm font-normal"
          >
            Sign Up
          </button>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>
