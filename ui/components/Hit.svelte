<script>
    import Cat from 'public/cat.svg'
    import { updatePlayer } from 'stores/updates'
    
    const handleClick = async () => await fetch('/api/v1/hit', { method: 'POST'})
        .then(r => r.json())
        .then(async data => {
            updatePlayer(data.player)
        })
</script>

<style>
    div {
        padding: 0 10rem;
        cursor: pointer;
    }
</style>

<div on:click={ handleClick }>
    {@html Cat}
</div>
