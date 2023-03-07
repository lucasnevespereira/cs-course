# Linked Lists

- Are comprised of nodes
- The nodes contain a link to the next node (and also the previous node for bidirectional linked lists)
- Can be unidirectional or bidirectional
- Are a basic data structure, and form the basis for many other data structures
- Have a single head node, which serves as the first node in the list
- Require some maintenance in order to add or remove nodes
- The methods we used are an example and depend on the exact use case and/or programming language being used

## Questions

 *TODO: Translate from Python to Go*

- What is the head node in a linked list?
  - Answer: The first node in the list

- Linked lists are made up of what elements?
  - Answer: Nodes

- Which of the following nodes is the head node of ll?
  ```
   class Node:
   def __init__(self, value, next_node=None):
    self.value = value
    self.next_node = next_node

    class LinkedList:
    def __init__(self, head_node=None):
     self.head_node = head_node
    
    a = Node(5)
    b = Node(70, a)
    c = Node(5675, b)
    d = Node(90, c)
    ll = LinkedList(d)
  ```
  - Answer: d is the head node of the linked list.

- How is a linked list terminated (in Python)?
  - Answer: By a node with a pointer to None

- Linked list nodes do not contain which of the following? (Roots, Links, Data, Pointers)
    - Answer: Roots

- What would you add to the .remove_node() method to properly maintain the linked list when removing a node?
  ```python
  class Node:
  def __init__(self, value, next_node=None):
    self.value = value
    self.next_node = next_node

    class LinkedList:
    def __init__(self, head_node=None):
    self.head_node = head_node
    
    def remove_node(self, node_to_remove):
    current_node = self.head_node
    if current_node == node_to_remove:
    self.head_node = current_node.next_node
    else:
    while current_node:
    next_node = current_node.next_node
    if next_node == node_to_remove:
    # --------> what line of code goes here?
    current_node = None
    else:
    current_node = next_node
  ```
  - Answer: currentNode.setNextNode(nextNode.getNextNode()). We can remove our next_node by setting our current node’s next_node property equal to the node that follows next_node.

- Consider the linked list: a -> b -> c. When removing b from the list, which node(s) need(s) to be updated?
  - Answer: a, only a needs to be updated to link to c.

- It is possible to traverse a linked list through its list property, which keeps track of each node in the list?
  - Answer: False. A linked list only keeps track of the first node in the list. To traverse the list, it needs a method that loops through each node to find the following node.

- Fix the Node class so that some_node = Node(6) can run without error.
  ```python
    class Node:
    def __init__(self, value, next_node):
     self.value = value
     self.next_node = next_node
    
    some_node = Node(6)
  ```
  
  - Answer: line 2 should be  `def __init__(self, value, next_node=None):`
    Setting a default for next_node would need to happen in the __init__ parameters. If another node is created that DOES pass an argument for next_node, setting next_node to None on line 4 would override that value

- What output would you expect to see in the terminal if you ran this code?
  ```python
    class Node:
    def __init__(self, value, next_node = None):
     self.value = value
     self.next_node = next_node
    
    class LinkedList:
     def __init__(self, head_node=None):
      self.head_node = head_node
    
     def stringify_list(self):
      string_list = ""
      current_node = self.head_node
      while current_node:
        string_list += str(current_node.value) + "."
        current_node = current_node.next_node
      return string_list
    
     a = Node(5)
     b = Node(70, a)
     c = Node(5675, b)
     d = Node(90, c)
     ll = LinkedList(d)
    
     print(ll.stringify_list())
  ```
  - Answer: 90.5675.70.5, because the first node in ll is d and we are traversing the list from beginning to end, we would expect the values printed in this order.

- Given this code, what would you add to complete the .add_new_head() method?
  ```python
  class Node:
   def __init__(self, value, next_node=None):
    self.value = value
    self.next_node = next_node

  class LinkedList:
    def __init__(self, head_node=None):
     self.head_node = head_node
    
    def add_new_head(self, new_head_node):
     # --------> what line of code goes here?
     self.head_node = new_head_node
  ``` 
  - Answer: new_head_node.next_node = self.head_node. It’s necessary to set the new head node’s next_node property equal to the current list’s head_node.
  