<script setup>
import { useQuasar } from 'quasar';
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const route = useRoute();

let session = route.query?.session || "";

// get the session from the http params
// save session as cookie
const $q = useQuasar();

$q.cookies.set('session', session);
$q.cookies.remove("state")

// redirect to the main page
router.push({ path: '/' });

</script>
