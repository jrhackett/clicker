<script>
    import Tile from 'components/Tile'
    import Hit from 'components/Hit'
    import Worth from 'components/Worth'
    import Nav from 'components/Nav'
    import colors from 'styles/colors'
    import { playerStore } from 'stores'
    import { updatePlayer, updateGeneratorsCost } from 'stores/updates'
    import { onMount } from 'svelte';
    
    let backgroundColor = colors.offWhite

    let player
    const unsubscribe = playerStore.subscribe(p => player = p);

    let userLoggedIn = false;
    window.onLoadCallback = () => {
        userLoggedIn = gapi.auth2.getAuthInstance().isSignedIn.get();
        if (userLoggedIn) {  
        // Handle login
        }
    }

    onMount(async () => await fetch('/api/v1/player')
        .then(r => r.json())
        .then(data => {
            updatePlayer(data.player)
            player.generators.forEach(async g => await updateGeneratorsCost(g.name))
        })
        .then(async () => setInterval(update, 200))
    )

    const update = async () => {
        if(player.generators.reduce((acc, g) => acc += g.count, 0) > 0)
            await fetch('/api/v1/player')
                .then(r => r.json())
                .then(data => updatePlayer(data.player))
    }

    const navHeight = '3rem'
</script>

<style>
	.app {
        background-color: var(--background-color);
        height: 100vh;
    }

    .container {
        display: flex;
        flex-direction: column;
        flex: 1;
        padding: 1rem 4rem;
        padding-top: var(--nav-height);
    }

    .inner {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
        flex-wrap: wrap;
        padding: 1rem 0;
    }

    ul {
        display: flex;
    }
</style>

<div class="app" style="--background-color:{ backgroundColor };">
    <Nav height={ navHeight } />
    <div class="container" style="--nav-height:{ navHeight };">
        <div class="inner">
            {#if !!player}
                <Worth />
                <Hit />
                <ul>
                    {#each player.generators as generator}
                        <Tile generator={ generator } />
                    {/each}
                </ul>
            {/if}
        </div>
    </div>
</div>
