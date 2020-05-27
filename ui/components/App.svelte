<script>
    import Home from 'components/Home'
    import Tile from 'components/Tile'
    import colors from 'styles/colors'
    import { onMount } from 'svelte';
    
    let backgroundColor = colors.offWhite

    let player;

    onMount(async () => {
        await load()
    })

    const load = async () => await fetch('/api/v1/player')
            .then(r => r.json())
            .then(data => {
                updatePlayer(data)
                player.generators.forEach(async g => await updateGeneratorsCost(g.name))
            })
            .then(async () => setInterval(update, 200))

    const update = async () => await fetch('/api/v1/player')
            .then(r => r.json())
            .then(data => updatePlayer(data))

    const handleClick = async name => {
        await fetch('/api/v1/buy', {
            method: 'POST',
            body: JSON.stringify({ name })
        })
            .then(r => r.json())
            .then(async data => {
                updatePlayer(data)
                await updateGeneratorsCost(player.generators.find(a => a.name == name).name)
            })
    }

    const updatePlayer = data => {
        player = {
            ...player,
            liquid: data.player.liquid ? data.player.liquid : player.liquid,
            generators: data.player.generators.map(g => ({
                ...(player ? player.generators.find(a => a.name == g.name) : {}),
                ...g
            }))
        }
    }

    const updateGeneratorsCost = async name => {
        await fetch(`api/v1/cost?name=${ name }&count=1`)
            .then(r => r.json())
            .then(data => {
                player = {
                    ...player,
                    generators: player.generators.map(g => {
                        if (g.name == name) {
                            return {
                                ...g,
                                cost: data.cost
                            }
                        }

                        return g
                    })
                }  
            })
    }
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
    }

    .inner {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
        flex-wrap: wrap;
    }

    ul {
        display: flex;
    }

    li {
        list-style: none;
        padding: 0.5rem 1rem;
    }
</style>

<div class="app" style="--background-color:{ backgroundColor };">
    <div class="container">
        <div class="inner">
            {#if player}
                <p>${ player.liquid.toFixed(0) }</p>
                <ul>
                    {#each player.generators as generator}
                        <li on:click={ () => handleClick(generator.name) }>
                            <Tile generator={ generator } />
                        </li>
                    {/each}
                </ul>
            {/if}
        </div>
    </div>
</div>
