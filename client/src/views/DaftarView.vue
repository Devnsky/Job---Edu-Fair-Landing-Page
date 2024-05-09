<script setup>
import { ref} from 'vue'
import ModalView from './Modal.vue'
import HeaderView from './Header.vue'
import FooterView from './Footer.vue'

const emit = defineEmits(['isModalOpen'])

const showModal = ref(false)

const isModalOpen = () => {
  showModal.value = !showModal.value
}
</script>
<template>
    <main class="flex flex-col min-h-screen">
    <header>
      <HeaderView />
    </header>
    <div v-if="showModal" class="flex-grow w-full flex items-center justify-center px-5 mx-auto bg-[url('../assets/image/bg.png')]">
        <div class="max-w-sm bg-white rounded-md overflow-hidden flex-1 z-20 border border-gray-200 p-7">
          <div class="p-2 flex justify-center">
            <img src="../assets/image/smkn7logo.png" alt="smk7logo" class="w-20">
          </div>
          <div class="py-2">
            <h1 class="text-center text-xl font-bold">Job & Edu Fair 2024</h1>
          </div>
          <hr>
          <div class="py-4">
            <h1 class="px-2 py-1.5 border border-blue-500 bg-blue-200 rounded pb-2 text-center text-xl font-semibold text-blue-700">Presensi Pengunjung</h1>
          </div>
          <hr>
          <div class="py-4 flex justify-center items-center">
            <button @click="isModalOpen" class="px-3 py-1.5 rounded bg-red-500 text-white font-semibold flex">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mt-1 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M14.293 3.293a1 1 0 0 1 1.414 0l1 1a1 1 0 0 1 0 1.414l-9 9a1 1 0 0 1-.39.242l-5 1a1 1 0 0 1-1.213-1.213l1-5a1 1 0 0 1 .242-.39l9-9zM15 5l-9 9-3-3 9-9 3 3z" clip-rule="evenodd" />
              </svg>
              Daftar
            </button>
          </div>
        </div>
    </div>
    <ModalView v-if="!showModal" @close="isModalOpen" />
    <footer class="bg-gray-900">
      <FooterView />
    </footer>
  </main>
</template>