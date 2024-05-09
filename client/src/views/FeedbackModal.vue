<script setup>
import HeaderView from './Header.vue'
import FooterView from './Footer.vue'
import { ref } from 'vue'
import { useFeedbackStore } from '@/stores/feedback';
import Swal from 'sweetalert2';

const feedbackStore = useFeedbackStore()
const errorMessage = ref('')
const form = ref({
    feedback: ''
})

const submitFeedback = async () => {
    if (form.value.feedback === '') {
        Swal.fire({
            icon: 'error',
            title: 'gagal',
            text: 'Feedback tidak boleh kosong!',
            showConfirmButton: false,
            timer: 1500
        })
    } else {
        try {
            let response = await feedbackStore.Register(form.value)
            errorMessage.value = response.message
            if (errorMessage.value === 'Success') {
                Swal.fire({
                    icon: 'success',
                    title: 'TerimaKasih!',
                    text: '',
                    showConfirmButton: false,
                    timer: 1500
                })
                form.value.feedback = ''
            } else {
                Swal.fire({
                    icon: 'error',
                    title: 'gagal',
                    text: errorMessage.value,
                    showConfirmButton: false,
                    timer: 1500
                })
            }
        } catch (error) {
            console.error(error)
        }
    }
}
</script>

<template>
    <main class="flex flex-col min-h-screen">
        <header>
            <HeaderView />
        </header>
        <div class="flex-grow w-full flex items-center justify-center px-5 mx-auto bg-[url('../assets/image/bg.png')]">
            <div class="w-screen overflow-hidden">
                <div class="flex items-center justify-center py-4 text-xl font-semibold">
                    <svg class="mt-2 mr-2 h-5 w-5" viewBox="0 0 18 18" xmlns="http://www.w3.org/2000/svg"
                        mirror-in-rtl="true" fill="currentColor">
                        <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
                        <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
                        <g id="SVGRepo_iconCarrier">
                            <path fill="currentColor"
                                d="M4.1 2h-.2A2.906 2.906 0 0 0 1 4.9v1.2A2.906 2.906 0 0 0 3.9 9h.2A2.906 2.906 0 0 0 7 6.1V4.9A2.906 2.906 0 0 0 4.1 2zM4 10a4.012 4.012 0 0 0-4 4v2.667a1.326 1.326 0 0 0 1.333 1.324l5.333.01A1.337 1.337 0 0 0 8 16.667V14a4.01 4.01 0 0 0-4-4zM18 1v4a1 1 0 0 1-1 1h-3.99l-1.19 1.88a.47.47 0 0 1-.32.12.538.538 0 0 1-.21-.05.493.493 0 0 1-.29-.45V6h-1a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h7a1 1 0 0 1 1 1z">
                            </path>
                        </g>
                    </svg>
                    <h1>Feedback</h1>
                </div>
                <div class="flex items-center justify-center">
                    <div class="bg-white px-5 py-5 rounded-md md:max-w-md">
                        <p class="text-sm font-semibold">Beri Kami Masukan</p>
                        <p class="text-xs">Agar Kami Bisa Lebih Baik Untuk Job & Edu Fair Tahun Depan!</p>
                        <textarea v-model="form.feedback"
                            class="px-1.5 py-1.5 border border-gray-400 w-full rounded-md my-2 text-xs"
                            placeholder="Kritik/Saran" cols="30" rows="10"></textarea>
                        <button @click="submitFeedback"
                            class="bg-blue-500 text-white px-3 py-1.5 rounded-md text-xs">Kirim</button>
                    </div>
                </div>
            </div>
        </div>
        <footer class="bg-gray-900">
            <FooterView />
        </footer>
    </main>
</template>