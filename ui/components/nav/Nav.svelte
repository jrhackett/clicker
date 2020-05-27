<script>
    import colors from 'styles/colors'
    import RouterLink from 'components/nav/RouterLink'

    export let name = 'Clicker'
    export let backgroundColor = colors.blue
    export let backgroundHoverColor = colors.lightBlue
    export let textColor = colors.offWhite
    export let navHeight = '3rem'

</script>

<style>
	header {
        position: fixed;
        width: 100%;
        height: var(--nav-height);
    }

    nav {
        display: flex;
        justify-content: space-between;
        padding: 0rem 2rem;
        background-color: var(--nav-background-color);
    }

    ul {
        display: block;
        list-style: none;
    }

    li {
        display: inline-flex;
        align-items: center;
        padding: 0 1rem;
        height: var(--nav-height);
        color: var(--text-color);
        cursor: pointer;
    }

    li:hover {
        background-color: var(--nav-background-hover-color);
    }
</style>

<header style="--nav-height:{ navHeight }; --text-color:{ textColor };">
	<nav style="--nav-background-color:{ backgroundColor }; --nav-background-hover-color:{ backgroundHoverColor };">
        <span>
            <RouterLink page={ { path: '/', name } } />
        </span>
        <ul>
            <li>
                <RouterLink page={ { path: '/', name: "Test" } } />
            </li>
        </ul>
	</nav>
</header>
