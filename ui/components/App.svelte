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
        display: flex;
        flex: 1;
        flex-direction: column;
        height: 100vh;
        background-color: var(--background-color);
    }

    .container {
        display: flex;
        flex-direction: column;
        flex: 1;
        padding-top: var(--nav-height);
        height: calc(100% - var(--nav-height));
    }

    .inner {
        display: flex;
        justify-content: space-between;
        height: 100%;
        flex-wrap: wrap;
    }

    .generators {
        display: flex;
        flex: 0.5;
        flex-direction: column;
        height: calc(100% - 10px);
        overflow-y: auto;
        border-right: 5px solid var(--border-color);
        border-top: 10px solid var(--border-color);
    }

    .hit {
        display: flex;
        flex: 0.5;
        flex-direction: column;
        height: calc(100% - 10px);
        justify-content: center;
        border-left: 5px solid var(--border-color);
        border-top: 10px solid var(--border-color);
        background-color: var(--background-color);
    }
</style>

<div class="app" style="--background-color:{ colors.orange };">
    <Nav height={ navHeight } />
    <div class="container" style="--nav-height:{ navHeight };">
        <div class="inner" style="--border-color:{ colors.grey }">
            {#if !!player}
                <div class="generators">
                    <ul>
                        {#each player.generators as generator}
                            <Tile generator={ generator } />
                        {/each}
                    </ul>
                </div>
                <div class="hit" style="--background-color:{ colors.yellow };">
                    <Worth />
                    <Hit />
                </div>
            {/if}
        </div>
    </div>
</div>
