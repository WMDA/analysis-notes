#Presentation specific settings
from traitlets.config.manager import BaseJSONConfigManager
from pathlib import Path
path = Path.home() / ".jupyter" / "nbconfig"
cm = BaseJSONConfigManager(config_dir=str(path))
cm.update(
    "rise",
    {
        "theme": "blood",
        "transition": "zoom",
        "start_slideshow_at": "selected",
         "overlay": "<div class='myheader'><h2>BEACON PROJECT</div><div class='myfooter'></div>",
         "enable_chalkboard": "true"
     }
)
