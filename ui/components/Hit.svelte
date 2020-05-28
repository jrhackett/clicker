<script>
    import { updatePlayer } from 'stores/updates'
    
    const handleClick = async () => await fetch('/api/v1/hit', { method: 'POST'})
        .then(r => r.json())
        .then(async data => {
            updatePlayer(data.player)
        })
</script>

<button on:click={ handleClick }>HIT!</button>
