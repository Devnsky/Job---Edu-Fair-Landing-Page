import axios from "axios";
import { defineStore } from "pinia";

export const useFeedbackStore = defineStore("Feedback", () => {
  function Register(form) {
    return new Promise(async (resolve, reject) => {
      try {
        let network = await axios.post(`api/feedback/register`, form);
        resolve(network.data);
      } catch (e) {
        reject(e);
      }
    });
  }
  return {
    Register,
  };
});
