<script>
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
</style>

<div class="container">
    <div class="offset">
        {#if player}
            <p>{ player.liquid }</p>
            <ul>
                {#each player.generators as generator}
                    <li>
                        <h3>{ generator.name }</h3>
                        <p>Cost: { generator.cost }</p>
                        <p>Count: { generator.count }</p>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
</div>
