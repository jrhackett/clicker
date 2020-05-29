<script>
    import LogIn from 'components/LogIn'
    import Logout from 'components/Logout'
    import Cat from 'public/cat.svg' 
    import colors from 'styles/colors'
    import { userStore } from 'stores'

    let user
    const unsubscribe = userStore.subscribe(u => user = u)

    export let height = '3rem'
</script>

<style>
	header {
        position: fixed;
        width: 100%;
        height: var(--height);
    }

    nav {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0rem 2rem;
        background-color: var(--nav-background-color);
        height: 100%;
    }

    div {
        display: flex;
        align-items: center;
        height: 100%;
        color: var(--text-color);
    }

    p {
        font-size: 1.5rem;
        padding-left: 0.5rem;
    }

    ul {
        display: block;
        list-style: none;
    }

    li {
        display: inline-flex;
        align-items: center;
        padding: 0 1rem;
        height: var(--height);
        color: var(--text-color);
        cursor: pointer;
    }

    li:hover {
        background-color: var(--nav-background-hover-color);
    }
</style>

<header style="--height:{ height };--text-color:{ colors.white };">
	<nav style="--nav-background-color:{ colors.blue }; --nav-background-hover-color:{ colors.lightGrey };">
        <div>
            {@html Cat}
            <p>Kittenfluencer</p>
        </div>
        <ul>
            {#if user.id !== 0}
                <li>{ user.name }</li>
                <Logout />
            {:else}
                <LogIn />
            {/if}
        </ul>
	</nav>
</header>
