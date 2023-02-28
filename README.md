# all the workspaces

Do more of them workspaces

### Install:

- download and install this repo (have to copy the `atws` into the global path, so somewhere into `/usr/bin/`)
- put `exec atws reset` somewhere at the begingin of your i3/config file
- paste
  ```
  # switch to workspace #workspace number $ws1
  bindsym $mod+1 exec atws workspace 1 
  bindsym $mod+2 exec atws workspace 2 
  bindsym $mod+3 exec atws workspace 3
  bindsym $mod+4 exec atws workspace 4
  bindsym $mod+5 exec atws workspace 5
  bindsym $mod+6 exec atws workspace 6
  bindsym $mod+7 exec atws workspace 7
  bindsym $mod+8 exec atws workspace 8
  bindsym $mod+9 exec atws workspace 9
  bindsym $mod+0 exec atws workspace 10
  
  bindsym $mod+Escape exec atws layer 0 
  bindsym $mod+F1 exec atws layer 1
  bindsym $mod+F2 exec atws layer 2
  bindsym $mod+F3 exec atws layer 3
  bindsym $mod+F4 exec atws layer 4
  bindsym $mod+F5 exec atws layer 5
  bindsym $mod+F6 exec atws layer 6
  bindsym $mod+F7 exec atws layer 7
  bindsym $mod+F8 exec atws layer 8
  bindsym $mod+F9 exec atws layer 9
  bindsym $mod+F10 exec atws layer 10
  bindsym $mod+F11 exec atws layer 11
  bindsym $mod+F12 exec atws layer 12
  
  # move focused container to workspace #move container to workspace number $ws1
  bindsym $mod+Shift+1 exec atws move-workspace 1
  bindsym $mod+Shift+2 exec atws move-workspace 2
  bindsym $mod+Shift+3 exec atws move-workspace 3
  bindsym $mod+Shift+4 exec atws move-workspace 4
  bindsym $mod+Shift+5 exec atws move-workspace 5
  bindsym $mod+Shift+6 exec atws move-workspace 6
  bindsym $mod+Shift+7 exec atws move-workspace 7
  bindsym $mod+Shift+8 exec atws move-workspace 8
  bindsym $mod+Shift+9 exec atws move-workspace 9
  bindsym $mod+Shift+0 exec atws move-workspace 0
  
  bindsym $mod+Shift+Escape exec atws move-layer 0 
  bindsym $mod+Shift+F1 exec atws move-layer 1
  bindsym $mod+Shift+F2 exec atws move-layer 2
  bindsym $mod+Shift+F3 exec atws move-layer 3
  bindsym $mod+Shift+F4 exec atws move-layer 4
  bindsym $mod+Shift+F5 exec atws move-layer 5
  bindsym $mod+Shift+F6 exec atws move-layer 6
  bindsym $mod+Shift+F7 exec atws move-layer 7
  bindsym $mod+Shift+F8 exec atws move-layer 8
  bindsym $mod+Shift+F9 exec atws move-layer 9
  bindsym $mod+Shift+F10 exec atws move-layer 10
  bindsym $mod+Shift+F11 exec atws move-layer 11
  bindsym $mod+Shift+F12 exec atws move-layer 12
  ```
  in your i3/config
