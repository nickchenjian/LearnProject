# 代理模式

 * 定义 ： 为其他对象提供一种代理以控制对这个对象的访问.
 * 代理模式的实现 一般由三个类组成： 抽象类，代理类，被代理对象
 * 个人认为，普通代理模式最精髓的一点就是，代理类能相互组合，灵活的定义不同的代理功能 比如例子中的 ProxtPlayer 和 ProxyLogger ，他们分别的做了嵌套，实现了paly 和 log 的功能        
