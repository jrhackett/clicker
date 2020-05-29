<script>
    import colors from 'styles/colors'
    import { updatePlayer, updateGeneratorsCost } from 'stores/updates'
    import { playerStore } from 'stores'

    export let generator = {}

    let player
    const unsubscribe = playerStore.subscribe(p => player = p);

    const handleClick = async name => await fetch('/api/v1/buy', {
        method: 'POST',
        body: JSON.stringify({ name })
    })
        .then(r => r.json())
        .then(async data => {
            updatePlayer(data.player)
            await updateGeneratorsCost(name)
        })
</script>

<style>
    li {
        list-style: none;
        flex-basis: 33.33%;
    }

    .inner {
        padding: 0.5rem;
    }

    .tile {
        padding: 0.5rem 1rem;
        border-radius: 2px; 
        color: var(--text-color);
        background-color: var(--tile-color);
    }
</style>

<li on:click={ () => handleClick(generator.name) }>
    <div class="inner">
        <div class="tile" style="--tile-color:{ player.liquid >= generator.cost ? colors.blue : colors.lightGrey };--text-color:{ colors.white }">
            <h3>{ generator.name }</h3>
            <p>Cost: { generator.cost && generator.cost.toFixed(0) } likes</p>
            <p>Owned: { generator.count }</p>
            <p>ROI: { generator.gained.toFixed(0) }</p>
            <p>Yield: { generator.gain_per_second }</p>
        </div>
    </div>
</li>
