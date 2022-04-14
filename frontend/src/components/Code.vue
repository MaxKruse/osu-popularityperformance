<script setup>
import { useQuasar } from 'quasar';
import { useRoute, useRouter } from 'vue-router';

const q = useQuasar();

const router = useRouter();
const route = useRoute();


// if we have a token, set the cookie now
let session = route.query?.session || "";

if (session !== "") {
    console.log("Session param is there:", session);
    q.cookies.set('session', session);
    q.cookies.remove("state");

    router.push({
        path: "/"
    });
}

</script>