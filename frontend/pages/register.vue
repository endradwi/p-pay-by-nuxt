<script lang="ts" setup>
useSeoMeta({
    title: 'Register Page'
})

import { ref } from 'vue';

const email = ref('');
const password = ref('');
const errorMessage = ref('');
const successMessage = ref('');

function register() {
  // Basic validation
  if (!email.value || !password.value) {
    errorMessage.value = 'Email and password are required.';
    return;
  }

  const value = {
    email: email.value,
    password: password.value,
  };

  const newData = new URLSearchParams(value);
  const nq = newData.toString();

  $fetch('http://localhost:8881/auth/register', {
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    method: 'POST',
    body: nq,
  })
  .then(response => {
    if (response.ok) {
      successMessage.value = 'Registration successful!';
      errorMessage.value = '';
    } else {
      errorMessage.value = 'Registration failed. Please try again.';
    }
  })
  .catch(error => {
    errorMessage.value = 'An error occurred: ' + error.message;
  });
}
</script>

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
          Transfering money is easier than ever, you can access Zwallet
          wherever you are. Desktop, laptop, mobile phone? we cover all of that
          for you!
        </div>
        <form @submit.prevent="register" class="w-full flex flex-col gap-5"> 
          <div v-if="errorMessage" class="text-red-500">{{ errorMessage }}</div>
          <div v-if="successMessage" class="text-green-500">{{ successMessage }}</div>

          <label htmlFor="" class="flex flex-col gap-1">
            <div>
              <span class="text-base font-medium">Email</span>
            </div>
            <div class="input border-info w-full h-11 pl-5 flex items-center gap-3">
              <MdOutlineMail class="text-info" />
              <input
                type="text"
                placeholder="Enter Your Email"
                v-model="email"
                class="w-full box-border"
              />
            </div>
          </label>
          <label htmlFor="" class="flex flex-col gap-1">
            <span class="text-base font-medium">Password</span>
            <div class="w-full h-11 border-black pl-5 flex items-center justify-between">
              <div class="flex items-center gap-3">
                <MdOutlineVpnKey class="text-info" />
                <input
                  type="text"
                  placeholder="Enter Your Password"
                  v-model="password"
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
