<script>
    import { userStore } from 'stores'
    import { updateUser } from 'stores/updates'

    let user
    const unsubscribe = userStore.subscribe(u => user = u)

    window.onSignIn = googleUser => {
        const profile = googleUser.getBasicProfile()
        updateUser(profile.getId(), profile.getName(), profile.getImageUrl(), profile.getEmail())
    }
</script>

<svelte:head>
    <script src="https://apis.google.com/js/platform.js?onload=onLoadCallback" async defer></script>
</svelte:head>

<div class="g-signin2" data-onsuccess="onSignIn" />

<p>{ user.name }</p>
