<script>
    import { playerStore } from 'stores'
    
    let player
    const unsubscribe = playerStore.subscribe(p => player = p);
</script>

<style>
    div {
        display: flex;
        align-items: center;
        flex-direction: column;
    }
</style>

<div>
    <p>${ player.liquid.toFixed(0) }</p>
    <p>${ player.generators.reduce((acc, g) => acc += g.count * g.gain_per_second, 0) }/second</p>
</div>
