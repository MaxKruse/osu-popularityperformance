import { defineStore } from 'pinia';

export const useStore = defineStore("main", {
    state: () => {
        return {
            user: {
                name: "",
                avatarUri: "",
                banchoId: 0
            }
        }
    }
})