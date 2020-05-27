<script>
    import colors from 'styles/colors'

    export let generator = {}
</script>

<style>
    .tile {
        padding: 0.5rem 1rem;
        background-color: var(--tile-color);
        border-radius: 2px; 
        color: var(--text-color);
    }
</style>

<div class="tile" style="--tile-color:{ colors.blue };--text-color:{ colors.white }">
    <h3>Name: { generator.name }</h3>
    <p>Cost: { generator.cost && generator.cost.toFixed(0) }</p>
    <p>Count: { generator.count }</p>
    <p>Gained: { generator.gained.toFixed(0) }</p>
    <p>Gain Per Second: { generator.gain_per_second }</p>
</div>
