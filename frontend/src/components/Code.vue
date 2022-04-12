<script setup>
import { inject, ref, toRef } from 'vue';
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const route = useRoute();

let session = route.query?.session || "";

console.log(session)
// get the session from the http params
// save session as cookie
const cookies = inject('$cookies');
cookies.set('session', session);
cookies.remove("state")

// redirect to the main page
router.push({ path: '/' });

</script>
