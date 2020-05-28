<script>
    import colors from 'styles/colors'
    import { updatePlayer, updateGeneratorsCost } from 'stores/updates'
    import { playerStore } from 'stores'

    export let generator = {}

    let player
    const unsubscribe = playerStore.subscribe(p => {
		player = p
	});

    const handleClick = async name => {
        await fetch('/api/v1/buy', {
            method: 'POST',
            body: JSON.stringify({ name })
        })
            .then(r => r.json())
            .then(async data => {
                updatePlayer(data.player)
                await updateGeneratorsCost(name)
            })
    }
</script>

<style>
    li {
        list-style: none;
        padding: 0.5rem 1rem;
    }

    .tile {
        padding: 0.5rem 1rem;
        background-color: var(--tile-color);
        border-radius: 2px; 
        color: var(--text-color);
    }
</style>

<li on:click={ () => handleClick(generator.name) }>
    <div class="tile" style="--tile-color:{ player.liquid >= generator.cost ? colors.blue : colors.lightGrey };--text-color:{ colors.white }">
        <h3>Name: { generator.name }</h3>
        <p>Cost: { generator.cost && generator.cost.toFixed(0) }</p>
        <p>Count: { generator.count }</p>
        <p>Gained: { generator.gained.toFixed(0) }</p>
        <p>Gain Per Second: { generator.gain_per_second }</p>
    </div>
</li>
