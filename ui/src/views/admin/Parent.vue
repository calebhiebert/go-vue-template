<template>
    <div>
        <div class="sidebar-page">
            <b-sidebar
                position="static"
                :fullheight="true"
                type="is-light" open>
                <div class="p-1">
                    <b-menu>
                        <b-menu-list label="Menu">
                            <b-menu-item icon="information-outline" label="Info"></b-menu-item>
                            <b-menu-item active expanded icon="settings" label="Data">
                                <b-menu-item icon="account" label="Users"></b-menu-item>
                            </b-menu-item>
                            <b-menu-item icon="account" label="My Account">
                                <b-menu-item icon="account-box" label="Account data"></b-menu-item>
                                <b-menu-item icon="home-account" label="Addresses"></b-menu-item>
                            </b-menu-item>
                        </b-menu-list>
                        <b-menu-list>
                            <b-menu-item label="Expo" icon="link"></b-menu-item>
                        </b-menu-list>
                        <b-menu-list label="Actions">
                            <b-menu-item icon="logout" label="Logout"></b-menu-item>
                        </b-menu-list>
                    </b-menu>
                </div>
            </b-sidebar>
            <div class="content">
                <router-view></router-view>
            </div>
        </div>

    </div>
</template>

<script>
export default {
    name: "Parent",

    mounted() {
        console.log("HEIY SODKIFSe");
    },
};
</script>

<style scoped>
.sidebar-page {
    display: flex;
    width: 100%;
    min-height: 100vh;
}

.content {
    width: 100%;
    margin-left: 1em;
}
</style>