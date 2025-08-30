const {FusesPlugin} = require('@electron-forge/plugin-fuses');
const {FuseV1Options, FuseVersion} = require('@electron/fuses');

module.exports = {
    packagerConfig: {
        asar: false,
        extraResource: [
            'dist',
            'Billadm-Kernel.exe'
        ]
    },
    rebuildConfig: {},
    makers: [
        {
            name: '@electron-forge/maker-squirrel',
            config: {},
        }
    ],
    plugins: [
        new FusesPlugin({
            version: FuseVersion.V1,
            [FuseV1Options.RunAsNode]: false,
            [FuseV1Options.EnableCookieEncryption]: true,
            [FuseV1Options.EnableNodeOptionsEnvironmentVariable]: false,
            [FuseV1Options.EnableNodeCliInspectArguments]: false,
            [FuseV1Options.EnableEmbeddedAsarIntegrityValidation]: true,
            [FuseV1Options.OnlyLoadAppFromAsar]: true,
        }),
    ],
};
