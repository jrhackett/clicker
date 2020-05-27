<script>
import colors from 'styles/colors'
    import { onMount } from "svelte";

    let player;

    onMount(async () => {
        await fetch(`/api/v1/player`)
            .then(r => r.json())
            .then(data => {
                console.log(data)
                player = data.player
            })
    })

    async function handleClick(name) {
        await fetch(`/api/v1/buy`, {
            method: 'POST',
            body: JSON.stringify({ name })
        })
            .then(r => r.json())
            .then(data => {
                console.log(data)
                player = data.player
            })
    }
</script>

<style>
    .container {
        display: flex;
        justify-content: center;
        align-items: center;
        height: 100%;
        flex-wrap: wrap;
    }

    .offset {
        margin-bottom: 20%;
    }

    ul {
        display: flex;
    }

    li {
        list-style: none;
        padding: 0.5rem 1rem;
    }

    .inner {
        padding: 0.5rem 1rem;
        background-color: var(--tile-color);
        border-radius: 2px; 
        color: var(--text-color);
    }
</style>

<div class="container">
    <div class="offset">
        {#if player}
            <p>${ player.liquid.toFixed(0) }</p>
            <ul>
                {#each player.generators as generator}
                    <li on:click={ () => handleClick(generator.name) }>
                        <div class="inner" style="--tile-color:{ colors.blue };--text-color:{ colors.white }">
                            <h3>Name: { generator.name }</h3>
                            <p>Cost: { generator.cost }</p>
                            <p>Count: { generator.count }</p>
                            <p>Gained: { generator.gained.toFixed(0) }</p>
                            <p>Gain Per Second: { generator.gain_per_second }</p>
                        </div>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
</div>
