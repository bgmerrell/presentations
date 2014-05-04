def get_value(self):
    return self.elem['value'] if not self.index else \
        self.elem['value'][self.index] if not 'marked' in self.elem['meta'] else \
        self.elem['value'][self.index][0]
