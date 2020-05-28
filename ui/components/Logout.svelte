<script>
    import { updateUser } from 'stores/updates'

    const signOut = () => gapi.auth2.getAuthInstance().signOut().then(() => updateUser(0, '', '', ''))
</script>

<button on:click={ signOut }>Sign out</button>
