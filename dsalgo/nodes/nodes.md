# Nodes

- Nodes contain data, which can be a variety of data types.
- They also contain links to other nodes. If a node has no links, or they are all null, you have reached the end of the path you were following.
- Nodes can be orphaned if there are no existing links to them.


# Questions

- Consider the following nodes and links: a -> n -> t. If you want to remove node n, but preserve node t, what should the resulting structure look like?
    - Answer: a -> t

- A node containing only null pointers indicates what?
  - Answer: You are at the end of the node path you were following

  - Which two features do most nodes contain?
    - Answer: Data and links to other nodes

- Which of the following methods implemented in the Node class are required to establish a Node class with an accessible but immutable value?
  ```
  type Node struct {
    value    string
    linkNode *Node
  }
  func (n *Node) getValue() string {
    return n.value
  }
  func (n *Node) getLinkNode() *Node {
    return n.linkNode
  }
    
  func (n *Node) setLinkNode(newLinkNode *Node) {
    n.linkNode = newLinkNode
  }
    
  func (n *Node) setValue(newValue string) {
    n.value = newValue
  }
      
  func (n *Node) appendValue(someValue string) {
    n.value = n.value + someValue
  }
  ```
  - Answer: getValue(), getLinkNode() and setLinkNode()

- Consider the following nodes and links: a -> n -> t. If you want to remove node n, but preserve node t, what are the steps you would take?
  - Answer: Change the link a to point to t using a.setLinkNode(t)