const demoModule = buildModule("DemoModule", (m) => {
    const { proxy, proxyAdmin } = m.useModule(proxyModule);

    const demo = m.contractAt("Demo", proxy);

    return { demo, proxy, proxyAdmin };
});
export default demoModule;